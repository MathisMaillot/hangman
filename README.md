# HangMan Game

Ce programme est une implémentation simple du jeu du Pendu en Go, utilisant uniquement des packages natifs.

## Prérequis

Assurez-vous d'avoir Go installé sur votre machine. Vous pouvez télécharger Go depuis [le site officiel de Go](https://golang.org/dl/).
Les fichiers pour jouer sont dans le répertoire mais vous pouvez importer votre popre dictionnaire dans le dossier main.

## Comment jouer

1. Clonez ce dépôt sur votre machine :

   ```bash
   git clone https://ytrack.learn.ynov.com/git/mmathis/hangman-classic.git

2. Accédez au répertoire du projet :

    ```bash
    cd hangman-classic/main

3. Exécutez le jeu :

    ```bash
    go run hangman.go *nom du dictionnaire*

Par exemple :

    go run hangman.go words.txt

/!\ Les mots dans le fichier doivent être en colone (comme dans words.txt) /!\


## Comment jouer

Un mot aléatoire va être tiré du dictionnaire et votre but sera de le devinez avant que la personne soit pendu (10 essaies).
Une lettre au hasard sera dévoilé afin de vous aider, 
/!\ Cela ne veut pas dire qu'elle y est qu'une seule fois /!\
Vous pouvez choisir une lettre ou plusieurs, si plusieur lettre ont été entré vous perdez 2 tentatives au lieu de une.

## Fonctionnalités

1. Pas de soucis si vous entrez deux fois la même lettre, ça ne comptera pas et vous ne perdez aucune tentative.

2. Vous pouvez écrire en majuscule ou en miniscule (ou même les deux en même temps)

3. Vous pouvez changer la police de la console en ASCII art, pour cela il faudra executer cette commande :

    ```bash
    go run hangman.go *nom du dictionnaire* --letterFile *nom du fichier contenant les caractères*

Par exemple :

    go run hangman.go words.txt --letterFile standard.txt

Vous pouvez importer vos caratères en importer un fichier avec les caractères "abcdefghijklmnopqrstuvwyxy_" dans cette ordre et disposés comme dans l'exemple standard.txt ou thinkertoy.txt

4. Vous pouvez vous arreter quand vous le voulez en tapant "stop" quand vous choisissez une lettre. Le jeu sera sauvegardé et vous pouvez le reprendre en lancer la commande :

    ```bash
    go run hangman.go --startWith save.txt


## Contributeurs

- MAILLOT Mathis (mathis.maillot@ynov.com)