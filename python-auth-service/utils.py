import jwt
import datetime

def generate_token(user_id):
    token = jwt.encode({'sub': user_id, 'exp': datetime.datetime.utcnow() + datetime.timedelta(hours=24)}, '1234', algorithm='HS256')
    return token
