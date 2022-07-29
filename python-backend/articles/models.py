from django.db import models


class Article(models.Model):
    title = models.CharField(null=True, blank=True, max_length=255)
    description = models.TextField(null=True)
