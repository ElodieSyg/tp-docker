Exercice 1 :

Plusieurs choses ne vont pas dans ce Dockerfile :
L’image de node a été importée mais aucune version n’est précisée
Le cache n’est pas optimisé, l’entièreté du code est mit en cache
Il s’agit d’un Dockerfile de développement, il adéquat pour la mise en production
L’accès est directement en root
Il manque un .dockerignore

Pour build mon image :
$ docker build -t tp-blog .

Pour lancer mon container :
$ docker run -d -p 80:80 --name blog-app-container --restart always tp-blog

Pour mettre mon image sur docker hub :
$ docker tag tp-blog elodiesyg/react-tp
$ docker push elodiesyg/react-tp

Exercice 3 :

Pour run mes containers :
$ docker-compose up
