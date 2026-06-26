# P1 
- Safety typing / condition (try/catch) - Crash x2
- Understand host and execution workflow
- Add a CLI /manager/
- Pooling fonctionnel + rety connect --> math/rand pour aléatoir sinon IOC trop simple
- Collecte d'informations sur l'hôte 
- HTTPclient complet + UA Légitime
- stockage côté serveur (full mémoire pour le moment)

# P2 
- Obfuscation & Log/comment removal
- TLS communication / Socket
- Web Manager (Flask)
- Obfuscation 
- Changement des routes /api/v1/users... 
- Mise en place des triggers command (die, host info, etc..)
- Encodage complet


# P3
- Persistence 
- Découpage selon type d'hôte 
- Ajout de sécurité connection
- Revoir tous les headers
- Base64 + AES + TLS - Payload chiffrement complet repos et transit



## Ideas 
Obsfucation 

```bash
go install mvdan.cc/garble@latest
garble build .  # au leu de go build

+ go build -ldflags="-s -w" . # Suprimer les symbole debug
```


