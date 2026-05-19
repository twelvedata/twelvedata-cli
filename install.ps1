#!/usr/bin/env pwsh
# Twelve Data CLI installer for Windows
#
# Usage (PowerShell):
#   irm https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.ps1 | iex
#
# Pin a version:
#   $env:TWELVEDATA_VERSION = 'v1.0.0'; irm https://raw.githubusercontent.com/twelvedata/twelvedata-cli/main/install.ps1 | iex
#
# Environment variables:
#   TWELVEDATA_INSTALL  - Custom install directory (default: $HOME\.twelvedata)
#   TWELVEDATA_VERSION  - Version to install (default: latest)

param(
  [string]$Version = $env:TWELVEDATA_VERSION
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

# --- Helpers -----------------------------------------------------------------

function Write-Info { param($msg) Write-Host "  $msg" -ForegroundColor DarkGray }
function Write-Ok   { param($msg) Write-Host "  $msg" -ForegroundColor Green }

function Write-Fail {
  param($msg)
  Write-Host "  error: $msg" -ForegroundColor Red
}

# --- Architecture detection --------------------------------------------------

if ($env:PROCESSOR_ARCHITECTURE -notin @('AMD64', 'EM64T')) {
  Write-Fail "Unsupported architecture: $env:PROCESSOR_ARCHITECTURE`n`n  Twelve Data CLI currently supports Windows x64 only."
  throw "Installation failed."
}

# --- Version + Download URL --------------------------------------------------

$repo = 'https://github.com/twelvedata/twelvedata-cli'
$target = 'windows-x64'

if ($Version) {
  $Version = $Version.TrimStart('v')
  if ($Version -notmatch '^\d+\.\d+\.\d+(-[a-zA-Z0-9.]+)?$') {
    Write-Fail "Invalid version format: $Version`n`n  Expected: semantic version like 1.0.0 or 1.2.3-beta.1`n  Usage:    `$env:TWELVEDATA_VERSION = 'v1.0.0'; irm <install-url> | iex"
    throw "Installation failed."
  }
  $url          = "$repo/releases/download/v$Version/twelvedata-$target.zip"
  $checksumsUrl = "$repo/releases/download/v$Version/checksums.txt"
} else {
  $url          = "$repo/releases/latest/download/twelvedata-$target.zip"
  $checksumsUrl = "$repo/releases/latest/download/checksums.txt"
}

$archiveName = "twelvedata-$target.zip"

# --- Install directory -------------------------------------------------------

if ($env:TWELVEDATA_INSTALL) { $installDir = $env:TWELVEDATA_INSTALL } else { $installDir = Join-Path $HOME '.twelvedata' }
$binDir     = Join-Path $installDir 'bin'
$exe        = Join-Path $binDir 'twelvedata.exe'

if (-not (Test-Path $binDir)) {
  New-Item -ItemType Directory -Path $binDir -Force | Out-Null
}

# --- Download + Extract ------------------------------------------------------

Write-Host ""
Write-Host "  Installing Twelve Data CLI..." -ForegroundColor White
Write-Host ""
Write-Info "Downloading from $url"
Write-Host ""

$tmpDir = Join-Path ([System.IO.Path]::GetTempPath()) "twelvedata-$([System.Guid]::NewGuid())"
New-Item -ItemType Directory -Path $tmpDir -Force | Out-Null
$tmpZip = Join-Path $tmpDir 'twelvedata.zip'

try {
  try {
    # Force TLS 1.2 for Windows PowerShell 5.1 (no-op on PowerShell 7+ where it is the default)
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12
    $ProgressPreference = 'SilentlyContinue'  # Invoke-WebRequest is ~10x faster without progress bar
    Invoke-WebRequest -Uri $url -OutFile $tmpZip -UseBasicParsing
  } catch {
    if ($Version) { $ver = $Version } else { $ver = 'latest' }
    Write-Fail "Download failed.`n`n  Possible causes:`n    - No internet connection`n    - The version does not exist: $ver`n    - GitHub is unreachable`n`n  URL: $url"
    throw "Installation failed."
  }

  # Verify SHA256 against the release's checksums.txt before extracting.
  $tmpChecksums = Join-Path $tmpDir 'checksums.txt'
  try {
    Invoke-WebRequest -Uri $checksumsUrl -OutFile $tmpChecksums -UseBasicParsing
  } catch {
    Write-Fail "Failed to download checksums.txt for integrity verification.`n`n  URL: $checksumsUrl"
    throw "Installation failed."
  }

  $expected = $null
  foreach ($line in Get-Content $tmpChecksums) {
    $parts = $line -split '\s+', 2
    if ($parts.Length -eq 2 -and $parts[1].Trim() -eq $archiveName) {
      $expected = $parts[0].Trim().ToLower()
      break
    }
  }
  if (-not $expected) {
    Write-Fail "checksums.txt has no entry for $archiveName -- refusing to install unverified binary."
    throw "Installation failed."
  }

  $actual = (Get-FileHash -Algorithm SHA256 -Path $tmpZip).Hash.ToLower()
  if ($actual -ne $expected) {
    Write-Fail "SHA256 mismatch -- refusing to install.`n`n  File:     $archiveName`n  Expected: $expected`n  Actual:   $actual"
    throw "Installation failed."
  }

  Write-Info "SHA256 verified"

  try {
    Expand-Archive -Path $tmpZip -DestinationPath $binDir -Force
  } catch {
    Write-Fail "Failed to extract archive: $_"
    throw "Installation failed."
  }
} finally {
  Remove-Item -Recurse -Force $tmpDir -ErrorAction SilentlyContinue
}

if (-not (Test-Path $exe)) {
  Write-Fail "Binary not found after extraction. The download may be corrupted -- try again."
  throw "Installation failed."
}

# --- Verify installation -----------------------------------------------------

try {
  $installedVersion = (& $exe --version 2>$null).Trim()
} catch {
  $installedVersion = 'unknown'
}

Write-Host ""
Write-Ok "Twelve Data CLI $installedVersion installed successfully!"
Write-Host ""
Write-Info "Binary:  $exe"

# --- PATH setup --------------------------------------------------------------

$userPath    = [Environment]::GetEnvironmentVariable('PATH', 'User')
if (-not $userPath) { $userPath = '' }
$pathEntries = $userPath -split ';' | Where-Object { $_ -ne '' }

if ($pathEntries -contains $binDir) {
  # Already on PATH -- just print the getting-started line
  Write-Host ""
  Write-Host "  Run " -NoNewline
  Write-Host "twelvedata --help" -ForegroundColor Cyan -NoNewline
  Write-Host " to get started"
  Write-Host ""
  Write-Info "Enable shell completion (optional):"
  Write-Host ""
  Write-Host "    twelvedata completion install" -ForegroundColor Cyan
  Write-Host ""
  return
}

# Add to user PATH (persists across sessions -- no admin rights needed)
$newPath = ($pathEntries + $binDir) -join ';'
[Environment]::SetEnvironmentVariable('PATH', $newPath, 'User')
$env:PATH = "$env:PATH;$binDir"  # Also update the current session

Write-Info "Added $binDir to PATH (User scope)"
Write-Host ""
Write-Info "Restart your terminal, then:"
Write-Host ""
Write-Info "Next steps:"
Write-Host ""
Write-Host "    `$env:TWELVEDATA_API_KEY = '...'" -ForegroundColor Cyan
Write-Host "    twelvedata --help" -ForegroundColor Cyan
Write-Host ""
Write-Info "Enable shell completion (optional):"
Write-Host ""
Write-Host "    twelvedata completion install" -ForegroundColor Cyan
Write-Host ""
return
