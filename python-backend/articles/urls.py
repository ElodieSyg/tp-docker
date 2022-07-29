from rest_framework import routers
from articles.views import ArticlesViewset


router = routers.DefaultRouter()

router.register("", ArticlesViewset, basename="articles")
urlpatterns = router.urls
