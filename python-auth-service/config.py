import os

class Config:
    SQLALCHEMY_DATABASE_URI = 'postgresql://postgres:1234@host.docker.internal:5432/jet-style'
    SQLALCHEMY_TRACK_MODIFICATIONS = False

class DevelopmentConfig(Config):
    DEBUG = True
