# build.ps1

$projectName = "GoEncrypt" 
$buildDir = "build"                                                
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
    & go build -o $outputFile "main.go"

    if ($LASTEXITCODE -eq 0) {
        Write-Host "Compilation reussie : $outputFile"
    } else {
        Write-Host "Erreur lors de la compilation."
        exit $LASTEXITCODE
    }
}

# Menu pour choisir l'action
Write-Host "Que souhaitez vous faire ?"
Write-Host "1. Construire le projet"
$choice = Read-Host "Entrez votre choix"

switch ($choice) {
    1 { BuildCode }
    default {
        Write-Host "Choix invalide. Veuillez entrer une proposition valide"
        exit 1
    }
}

Write-Host "Processus termine."
