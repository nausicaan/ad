package gizmo

var province = []string{"AB", "BC", "MB", "NB", "NC", "NL", "NS", "ON", "PE", "QC", "SK"}

// A list of CFIA domain controllers.
var cfia = []string{"CFONK1AWPDCP004", "CFABT2EWPDCP004", "CFABT2LWPDCP002", "CFBCV5CWPDCP002", "CFMBR3EWPDCP002", "CFNBE1CWPDCP004", "CFNSB3BWPDCP002", "CFONK1AWPDCP004", "CFONK1AWVDCP007", "CFONK1AWVDCP008", "CFONL5TWPDCP002", "CFONM3JWPDCP002", "CFONN1GWPDCP002", "CFONN1HWPDCP002", "CFQCH3AWPDCP002", "CFQCJ2SWPDCP002", "CFSKS7NWPDCP002"}

// An array with 177 rows and 2 columns to hold the bilingual text options.
var language = [181][2]string{
	{"for Active Directory ...and more!", "Pour Active directory ...et plus!"},
	{"Go language version created by:", "Version Go cree par:"},
	{"Testing and support:", "Test et support:"},
	{"WARNING!!! You aren't logged in with your ADM account", "ATTENTION!!! Vous n'etes pas connecte avec votre compte ADM"},
	{"Welcome", "Bienvenue"},
	{"Please make a selection: ", "S'il vous plait faire une selection: "},
	{"RESET Password", "REINITIALISER un mot de passe"},
	{"LOCKED OUT Account", "DEVEROUILLER un compte"},
	{"USER Information", "Information UTILISATEUR"},
	{"COMPUTER Information", "Information de l'ORDINATEUR"},
	{"PRINTER Information", "Information IMPRIMANTE"},
	{"GROUP Information", "Information GROUPE"},
	{"Advanced Computer Tools", "Outils Avances pour Ordinateur"},
	{"Back", "Dos"},
	{"Exit", "Sortie"},
	{"Advanced tools are currently linked to", "Outils avances est presentement connecte a"},
	{"Connectivity", "Connectivite"},
	{"Connected", "Connecte"},
	{"Advanced Tools", "Outils Avances"},
	{"Force Logoff", "Forcer la deconnection de l'usager"},
	{"Restart Computer", "Redemarrage de l'ordinateur"},
	{"Test Network Connection", "Tester la connection reseau"},
	{"Disable Network Card", "Desactiver la carte reseau"},
	{"Process Tools", "Outils pour les processus"},
	{"Service Tools", "Outils pour les services"},
	{"Event Logs", "Acceder aux journaux d'evenement"},
	{"Connection failed. Computer unreachable", "La connection a echouee. L'ordinateur ne peut etre rejoint"},
	{"Press Enter to continue...", "Appuyez sur retour pour continuer..."},
	{"Get Process List", "Afficher les processus"},
	{"Terminate Process", "Arreter un/des processus"},
	{"Choose an operation: ", "Choisissez une operation: "},
	{"List of processes running on", "Liste des processus en cours sur"},
	{"Enter the process name that matches the one(s) you want to terminate", "Entrez le nom du processus que vous"},
	{"No match for", "Pas de correspondance pour"},
	{"Are you sure you want to terminate this/those process(es)", "Etes-vous certain de vouloir arreter ce/ces"},
	{"Starting kill commands", "Commencement de/des arret(s)"},
	{"Killing", "Arret"},
	{"Killing completed", "Arret complete"},
	{"Kill command canceled", "Commande d'arret annulee"},
	{"Get Service List", "Afficher les services"},
	{"Start Service(s)", "Demarrer un/des service(s)"},
	{"Restart Service(s)", "Redemarrer un/des service(s)"},
	{"Stop Service(s)", "Arreter un/des service(s)"},
	{"List of services running on", "Liste des services sur"},
	{"Enter the service name(s) that match the one(s) you want to start", "Entez le nom du/des service(s) a demarrer"},
	{"Are you sure you want to start this/those service(s)", "Etes-vous certain de vouloir demarrer ce(s) service(s)"},
	{"Starting service(s)", "Demarrage du/des service(s)"},
	{"Starting", "Demarrage de"},
	{"Operation completed", "Operation completee"},
	{"Operation canceled", "Operation annulee"},
	{"Enter the service name(s) that match the one(s) you want to restart", "Entrez le nom du/des service(s) a"},
	{"Are you sure you want to restart this/those service(s)", "Etes-vous certain de vouloir redemarrer ce(s)"},
	{"Restarting service(s)", "Redemarrage de(s) service(s)"},
	{"Restarting", "Redemarrage de"},
	{"Enter the service name(s) that match the one(s) you want to stop", "Entrez le nom du/des service(s) a arreter"},
	{"Are you sure you want to stop this/those service(s)", "Etes-vous certain de vouloir arreter ce(s) service(s)"},
	{"Stoping service(s)", "Arret du/des service(s)"},
	{"Stopping", "Arret de"},
	{"SELECT DOMAIN(S)", "CHOISIR UN DOMAINE(S)"},
	{"Please select a Domain", "SVP choisir un domaine"},
	{"All", "Tous les domaines"},
	{"CFIA (Canadian Food Inspection Agency)", "ACIA (Agence Canadienne d'Inspection des Aliments)"},
	{"AAFC (Agriculture and Agri-Food Canada)", "AAC (Agriculture et Agroalimentaire Canada)"},
	{"Cancel", "Annuler"},
	{"USER SEARCH", "RECHERCHE D'USAGER"},
	{"Enter a USERNAME: ", "Entrez un NOM D'USAGER: "},
	{"Please wait", "Veuillez patienter"},
	{"No user found", "Aucun usager trouve"},
	{"Enter a USERNAME or Press Enter to exit: ", "Entrez un NOM D'USAGER ou appuyez sur Retour pour sortir: "},
	{"Listing found users", "Usager(s) trouve(s)"},
	{"COMPUTER SEARCH", "RECHERCHE D'ORDINATEUR"},
	{"COMPUTERNAME or IP Address: ", "NOM DE L'ORDINATEUR ou ADRESSE IP: "},
	{"No computer found", "Aucun ordinateur trouve"},
	{"Enter a COMPUTERNAME or Press Enter to exit: ", "Entrez un NOM D'ORDINATEUR ou Retour pour sortir: "},
	{"PRINTER SEARCH", "RECHERCHE D'IMPRIMANTE"},
	{"Printer NAME: ", "NOM d'imprimante: "},
	{"No printer found", "Aucune imprimante trouvee"},
	{"GROUP SEARCH", "RECHERCHE DE GROUPE"},
	{"Group name: ", "NOM de groupe: "},
	{"No group found", "Aucun groupe trouve"},
	{"Ok lets reset", "Ok reinitialison"},
	{"Enter your NEW PASSWORD: ", "Entrer le NOUVEAU MOT DE PASSE: "},
	{"Confirm your NEW PASSWORD: ", "Confirmer le NOUVEAU MOT DE PASSE: "},
	{"The passwords aren't the same, try again", "Les mots de passe sont differents, reessayez"},
	{"Change password at next logon", "Changer le mot de passe a la prochaine connection"},
	{"Let's go", "Allons-y"},
	{"ERROR! Something didn't work", "Une erreur est survenue"},
	{"MAKE SURE THE PASSWORD RESPECTS THE MINIMUM REQUIERMENTS", "ASSUREZ VOUS QUE LE MOT DE PASSE RESPECT LES CARACTERISTIQUES MINIMUMS"},
	{"The password was successfully reset", "Le mot de passe a ete change avec succes"},
	{"Do you want to check if the account is locked out", "Voulez-vous verifier si le compte est verouille"},
	{"Retreiving user information on multiple Domain Controllers. Please wait...", "Recuperation de l'information de l'usager sur de multiple controleur de domaine. SVP patienter..."},
	{"The account", "Le compte"},
	{"doesn't appear to be locked.", "ne semble pas etre verouille."},
	{"account is locked out.", "Ce compte est verouille."},
	{"Do you want to unlock it now? (Y\\N) ", "Est-ce que vous voulez le deverouiller maintenant? (Y\\N) "},
	{"Unlocking", "Deverouillage"},
	{"Unlocking account on", "Deverouillage en cours sur"},
	{"I think I unlocked", "J'ai deverouille le compte"},
	{"You can try to logon now", "Vous pouvez essayer de vous connecter"},
	{"The account is disabled", "Le compte est desactive"},
	{"The account is locked out", "Le compte est verouille"},
	{"The account is expired", "Le compte est expire"},
	{"The password is expired", "Le mot de passe est expire"},
	{"The user must change password", "L'usager doit changer son mot de passe"},
	{"Account status", "Etat du compte"},
	{"Quick summary for:", "Sommaire rapide pour:"},
	{"Computer name", "Nom de l'ordinateur"},
	{"Username", "Nom d'usager"},
	{"Full Name", "Nom complet"},
	{"Last time computer was rebooted/started from off", "Dernier redemarrage complet"},
	{"No user seems to be logged on at the moment", "Aucun usager ne semble etre connecte presentement"},
	{"Last boot up time", "Temps du dernier demarrage"},
	{"Quick summary not available for this computer at the moment", "Sommaire rapide indisponible pour le moment"},
	{"FULL COMPUTER DETAILS", "DETAIL COMPLET DE L'ORDINATEUR"},
	{"Testing connection to", "Test de la connection vers"},
	{"BIOS Information", "Information du BIOS"},
	{"MOTHERBOARD Information", "Information de la CARTE MERE"},
	{"OPERATING SYSTEM Information", "Information du SYSTEME D'OPERATION"},
	{"ENVIRONMENT VARIABLES Information", "Information des VARIABLES D'ENVIRONEMENT"},
	{"PROCESSOR(s) Information", "Information des PROCESSEURS"},
	{"MEMORY Information", "Information de la MEMOIRE"},
	{"HARD DRIVE(s) Information", "Information des DISQUES"},
	{"NETWORK CARD(s) Information", "Information des CARTES RESEAUX"},
	{"VIDEO CARD(s) Information", "Information des CARTES VIDEOS"},
	{"PAGE FILE information", "Fichier de pagination"},
	{"This computer is not responding", "Cet ordinateur ne repond pas"},
	{"LIST INSTALLED SOFTWARE", "LISTER LES LOGICIELS INSTALLES"},
	{"PRODUCT Information ... (A new window will open with the result.)", "Information des PRODUITS ...(Une nouvelle fenetre s'ouvrira avec le resultat.)"},
	{"List of software installed on", "Liste des logiciels installes sur"},
	{"LAUNCH ADVANCED TOOLS", "Lancer les OUTILS AVANCES"},
	{"The following information was found for:", "L'information suivante a ete trouvee pour:"},
	{"The printer is responding", "L'imprimante repond"},
	{"The printer is not responding", "L'imprimante ne repond pas"},
	{"cannot be verified", "ne peut pas etre verifie"},
	{"General Information", "Information generale"},
	{"Members (Name	 OU)", "Membre (nom	 OU)"},
	{"Nested Members (Name	 OU)", "Membres sous-jacents (nom	 OU)"},
	{"Warning: This selection is not allowed for the Printer list function ... Please go back and select a single domain", "Attention: Cette selection n'est pas disponible pour les listes complete d'imprimante ... Retournez en arriere et selectionnez un seul domaine."},
	{"Retrieving printer list on", "Recherche d'imprimante sur"},
	{"Printer List", "Liste d'imprimante"},
	{"Advanced tools are currently linked to", "Outils avances sont presentement connectes a"},
	{"WARNING!!! YOU ARE ABOUT TO LOGOFF ALL USERS USING", "ATTENTION!!! VOUS ETES SUR LE POINT DE DECONNECTER TOUS LES USAGERS UTILISANT"},
	{"Enter the word LOGOFF to proceed or press Enter to cancel the operation", "Entrez le mot LOGOFF pour proceder ou appuyer sur Retour pour annuler l'operation"},
	{"The logoff command has been sent to", "La commande de deconnection a ete envoyee a"},
	{"The logoff command has been canceled", "La commande de deconnection a ete annulee"},
	{"WARNING!!! YOU ARE ABOUT TO RESTART", "ATTENTION!!! VOUS ETES SUR LE POINT DE REDEMARRER"},
	{"Enter the word RESTART to proceed or press Enter to cancel the operation", "Entrez le mot RESTART pour proceder ou appuyer sur Retour pour annuler l'operation"},
	{"The restart command has been sent to", "La commande de redemarrage a ete envoyee a"},
	{"The restart command has been canceled", "La commande de redemarrage a ete annulee"},
	{"Checking connection to", "Verification de la connection vers"},
	{"Connection succeeded!", "Connection etablie!"},
	{"Connection failed!", "Connection a ??chou??!"},
	{"Testing speed", "Test de la vitesse"},
	{"Trying DOS ping instead ... (may generate errors)", "Utilisation du PING de DOS ... (peut generer des erreurs)"},
	{"Enter the DEVICE ID # of the card you want to disable", "Entrez le numero de 'DEVICE ID' de la carte que vous voulez desactiver"},
	{"WARNING!!! YOU ARE ABOUT TO DISABLE THE FOLLOWING CARD ON", "ATTENTION!!! VOUS ETES SUR LE POINT DE DESACTIVER LA CARTE SUIVANTE"},
	{"Enter the word DISABLE to proceed or press Enter to cancel the operation", "Entrez le mot DISABLE pour proceder ou appuyer sur Retour pour annuler l'operation"},
	{"The Disable Card command has been sent to", "La commande de desactivation a ete envoyee a"},
	{"The disable command has been canceled", "La commande de desactivation a ete annulee"},
	{"ORCA Status Verification", "Verificateur ORCA"},
	{"ORCA status checker", "Verificateur ORCA"},
	{"Enter the family name", "Entrez le nom de famille"},
	{"Enter the given (a.k.a first) name", "Entrez le prenom"},
	{"Search in progress. Please be patient while results are returned", "Recherche en cours. S'il vous plait etre patient tandis que les resultats sont retournes"},
	{"No users found. Ensure you have the correct spelling of the names", "Aucun utilisateur trouve. Assurez-vous d'avoir la bonne orthographe des noms"},
	{"Search for another user?", "Rechercher un autre utilisateur?"},
	{"Return to main menu", "Retour au menu principal"},
	{"IS an ORCA user", "EST un utilisateur ORCA"},
	{"is NOT an ORCA user", "n'est PAS un utilisateur ORCA"},
	{"You MUST enter at least one character for the family name otherwise too many results are returned", "Vous devez entrer au moins un caractere pour le nom de famille sinon trop the resultats sont retournes"},
	{"Y- Search another user", "O- Rechercher un autre utilisateur"},
	{"Full Name Search", "Recherche par nom complet"},
	{"Email Search", "Recherche par Courriel"},
	{"PowerShell version created by:", "Version originale de PowerShell cree par:"},
	{"Main Menu", "Menu Principal"},
	{"Region", "R??gion"},
	{"Locked", "Verouille"},
	{"Danger", "Danger"},
	{"Good", "Bon"},
	{"Response Time:", "Temps de r??ponse:"},
	{"Average Response Time:", "Temps de r??ponse moyen:"},
}
