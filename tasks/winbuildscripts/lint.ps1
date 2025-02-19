$ErrorActionPreference = "Stop"

$Env:Python3_ROOT_DIR=$Env:TEST_EMBEDDED_PY3

if ($Env:TARGET_ARCH -eq "x64") {
    & ridk enable
}
& $Env:Python3_ROOT_DIR\python.exe -m  pip install -r requirements.txt

$LINT_ROOT=(Get-Location).Path
$Env:PATH="$LINT_ROOT\dev\lib;$Env:GOPATH\bin;$Env:Python3_ROOT_DIR;$Env:Python3_ROOT_DIR\Scripts;$Env:PATH"

& $Env:Python3_ROOT_DIR\python.exe -m pip install PyYAML==5.3.1

$archflag = "x64"
if ($Env:TARGET_ARCH -eq "x86") {
    $archflag = "x86"
}

& inv -e deps

& inv -e rtloader.make --python-runtimes="$Env:PY_RUNTIMES" --install-prefix=$LINT_ROOT\dev --cmake-options='-G \"Unix Makefiles\"' --arch $archflag
$err = $LASTEXITCODE
Write-Host Build result is $err
if($err -ne 0){
    Write-Host -ForegroundColor Red "rtloader make failed $err"
    [Environment]::Exit($err)
}

& inv -e rtloader.install
$err = $LASTEXITCODE
Write-Host rtloader install result is $err
if($err -ne 0){
    Write-Host -ForegroundColor Red "rtloader install failed $err"
    [Environment]::Exit($err)
}

& inv -e install-tools
& inv -e lint-go --arch $archflag

$err = $LASTEXITCODE
Write-Host Lint result is $err
if($err -ne 0){
    Write-Host -ForegroundColor Red "lint failed $err"
    [Environment]::Exit($err)
}
