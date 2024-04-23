from flask_restful import Resource, reqparse
from __init__ import db
from models import User
from utils import generate_token
from flask_jwt_extended import jwt_required, get_jwt_identity
from sqlalchemy import text
from sqlalchemy.exc import SQLAlchemyError

class RegisterUser(Resource):
    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument('username', type=str, required=True, help='Username is required')
        parser.add_argument('password', type=str, required=True, help='Password is required')
        args = parser.parse_args()

        username = args['username']
        password = args['password']
        
        if User.query.filter_by(username=username).first():
            return {'message': 'Username already exists'}, 400

        user = User(username=username, password=password)
        db.session.add(user)
        db.session.commit()

        token = generate_token(user.id)

        return {'message': 'User registered successfully', 'token': token}, 201

class LoginUser(Resource):
    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument('username', type=str, required=True, help='Username is required')
        parser.add_argument('password', type=str, required=True, help='Password is required')
        args = parser.parse_args()

        username = args['username']
        password = args['password']

        user = User.query.filter_by(username=username, password=password).first()

        if not user:
            return {'message': 'Invalid username or password'}, 401

        token = generate_token(user.id)

        return {'message': 'User logged in successfully', 'token': token}, 200
    
class GetMyInfo(Resource):
    @jwt_required()
    def get(self):
        user_id = get_jwt_identity()

        sql_query = text("""
            SELECT id, username, registered_at, last_password_change
            FROM users
            WHERE id = :user_id
        """)

        try:
            user = db.session.execute(sql_query, {"user_id": user_id}).fetchone()

            if not user:
                return {'message': 'User not found'}, 404
            
            user_info_dict = {
                'id': user[0],
                'username': user[1],
                'registered_at': user[2].isoformat(),
                'last_password_change': user[3].isoformat() if user[3] else None
            }

            return user_info_dict, 200
        except SQLAlchemyError as e:
            return {'message': str(e)}, 500

class ChangeMyName(Resource):
    @jwt_required()
    def patch(self):
        parser = reqparse.RequestParser()
        parser.add_argument('new_username', type=str, required=True, help='New username is required')
        args = parser.parse_args()
        new_username = args['new_username']
        user_id = get_jwt_identity()

        try:
            sql_query = text("UPDATE users SET username = :new_username WHERE id = :user_id RETURNING id")
            result = db.session.execute(sql_query, {'new_username': new_username, 'user_id': user_id})
            updated_user_id = result.fetchone()[0]
            db.session.commit()
        except Exception as e:
            db.session.rollback()
            return {'message': str(e)}, 500

        if not updated_user_id:
            return {'message': 'User not found'}, 404

        return {'message': 'Username changed successfully'}, 200

class ChangeMyPassword(Resource):
    @jwt_required()
    def patch(self):
        parser = reqparse.RequestParser()
        parser.add_argument('new_password', type=str, required=True, help='New username is required')
        args = parser.parse_args()
        new_password = args['new_password']
        user_id = get_jwt_identity()

        try:
            sql_query = text("UPDATE users SET password = :new_password WHERE id = :user_id RETURNING id")
            result = db.session.execute(sql_query, {'new_password': new_password, 'user_id': user_id})
            updated_user_id = result.fetchone()[0]
            db.session.commit()
        except Exception as e:
            db.session.rollback()
            return {'message': str(e)}, 500

        if not updated_user_id:
            return {'message': 'User not found'}, 404

        return {'message': 'Password changed successfully'}, 200

