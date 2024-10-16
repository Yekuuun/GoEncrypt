# build.ps1

$projectName = "GoEncrypt" 
$buildDir = "build"                          
$sourceDir = "cmd"                          
$outputFile = "$buildDir\$projectName.exe" 

# Verifier si Go est install
if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Go n'est pas installe. Veuillez installer Go avant de continuer."
    exit 1
}

# Créer le répertoire de build s'il n'existe pas
if (-not (Test-Path $buildDir)) {
    New-Item -ItemType Directory -Path $buildDir
}

# Fonction pour construire le projet
function BuildCode {
    Write-Host "Compilation de l'application..."
    & go build -o $outputFile "$sourceDir\main.go"

    if ($LASTEXITCODE -eq 0) {
        Write-Host "Compilation reussie : $outputFile"
    } else {
        Write-Host "Erreur lors de la compilation."
        exit $LASTEXITCODE
    }
}

# Fonction pour exécuter les tests
function RunTests {
    Write-Host "Execution des tests..."
    & go test ./...

    if ($LASTEXITCODE -eq 0) {
        Write-Host "Tous les tests ont reussi."
    } else {
        Write-Host "Des tests ont echoue."
        exit $LASTEXITCODE
    }
}

# Menu pour choisir l'action
Write-Host "Que souhaitez vous faire ?"
Write-Host "1. Construire le projet"
Write-Host "2. Executer les tests"
Write-Host "3. Construire le projet et executer les tests"
$choice = Read-Host "Entrez votre choix (1/2/3)"

switch ($choice) {
    1 { BuildCode }
    2 { RunTests }
    3 { 
        BuildCode 
        RunTests 
    }
    default {
        Write-Host "Choix invalide. Veuillez entrer 1, 2 ou 3."
        exit 1
    }
}

Write-Host "Processus termine."
