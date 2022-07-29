from rest_framework.viewsets import ModelViewSet
from articles.models import Article
from articles.serializers import ArticleSerializer


class ArticlesViewset(ModelViewSet):
    queryset = Article.objects.all()
    serializer_class = ArticleSerializer
