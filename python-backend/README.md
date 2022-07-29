# Installation du projet

## Développement local

- Installation de l'environnement virtuel

```sh
python -m venv venv
./venv/bin/activate # Mac / Linux
.\venv\Scripts\Activate.ps #Windows
pip install -r requirements.tx
```

- Lancement du projet
```sh
(venv) python manage.py migrate
(venv) python manage.py runserver
```

## Variables de configuration

- **DATABASE_NAME**: Le nom de la base (Par défaut: `articles_database`)
- **DATABASE_USER**: L'utilisateur pour se connecter à la base (Par défaut: `articles`)
- **DATABASE_PASSWORD**: Le mot de passe de l'utilisateur (Par défaut: `fXcderT543--T(43`)
- **DATABASE_HOST**: L'adresse de la base de données (Par défaut: `localhost`)
- **DATABASE_PORT**: Le port de connexion à la base (Par défaut: `3306`)
