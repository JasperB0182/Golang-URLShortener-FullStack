# Backend URL Shortener

Deze URL shortener is gebouwd in **Go (Golang)**, de programmeertaal ontwikkeld door Google.  
Het doel van deze documentatie is om duidelijk uit te leggen wat elk bestand en elke map in deze repository doet, zodat je eenvoudig de structuur en werking van het project kunt begrijpen.

---

## Controllers

### **AdminController.go**
Bevat alle functionaliteiten die voor admins bedoeld zijn, zoals het inzien en beheren van alle URL’s.

### **creditController.go**
Functies om credits toe te voegen aan een account. Hiermee kun je testen of URL-aankopen werken en of een onbeperkte einddatum geselecteerd kan worden.

### **shortenerController.go**
Alle functionaliteiten rondom het verkorten van URLs, het ophalen ervan, het uitschakelen van URL’s en het opnieuw activeren.

### **usersController.go**
Functies voor gebruikersbeheer, zoals accounts aanmaken, inloggen en het wijzigen van accountgegevens.

---

## Initializers

### **connectToDb.go**
Verbindt de API met de database. Deze verbinding wordt door het hele project gebruikt.

### **loadEnvVariables.go**
Laadt alle environment variables, zoals database-inloggegevens.

### **syncDatabase.go**
Wordt bij het starten van de API uitgevoerd. Hiermee worden alle models gesynchroniseerd met de database.

---

## Middleware

### **rateLimiter.go**
Beperkt het aantal gelijktijdige requests om te voorkomen dat het systeem overbelast raakt.

### **requireAdmin.go**
Controleert of de gebruiker die het request verstuurt adminrechten heeft.

### **requireAuth.go**
Controleert of de gebruiker die het request verstuurt ingelogd is.

---

## Models

### **linkModel.go**
Model voor URL’s en hoe deze in de database worden opgeslagen.

### **roleModel.go**
Model voor gebruikersrollen en hoe deze worden opgeslagen.

### **userModel.go**
Model voor gebruikers en hun gegevens in de database.

---

## Seeders

### **databaseSeeder.go**
Bij het opstarten van het programma worden standaardrollen automatisch aangemaakt in de database.

---

## main.go

Het hoofdprogramma dat moet worden uitgevoerd om de API te starten.  
Na het opstarten draait de server standaard op:

---

## .env file setup

PORT=8?
DATABASE_URL=?
JWT_SECRET=?
BASE_URL=?

---

## Hoe het project op te zetten

git clone https://gitlab.inf-hsleiden.nl/s1155772/kecoex_s1155772.git
cd Backend
