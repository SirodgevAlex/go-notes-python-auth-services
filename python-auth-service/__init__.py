from flask import Flask
from flask_restful import Api
from flask_sqlalchemy import SQLAlchemy
from flask_jwt_extended import JWTManager
from config import DevelopmentConfig

app = Flask(__name__)
app.config.from_object(DevelopmentConfig)

api = Api(app)

db = SQLAlchemy(app)

app.config['JWT_TOKEN_LOCATION'] = ['headers']
app.config['JWT_SECRET_KEY'] = '1234'
jwt = JWTManager(app)

from resources import RegisterUser, LoginUser, GetMyInfo, ChangeMyName, ChangeMyPassword

api.add_resource(RegisterUser, '/register')
api.add_resource(LoginUser, '/login')
api.add_resource(GetMyInfo, '/me')
api.add_resource(ChangeMyName, '/me')
api.add_resource(ChangeMyPassword, '/me/password')

if __name__ == '__main__':
    app.run()
