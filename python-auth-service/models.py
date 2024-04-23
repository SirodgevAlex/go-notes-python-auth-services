import re
from __init__ import db
from sqlalchemy.sql import func

class User(db.Model):
    __tablename__ = 'users'

    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(50), unique=True, nullable=False)
    password = db.Column(db.String(255), nullable=False)
    registered_at = db.Column(db.DateTime(timezone=True), server_default=func.now())
    last_password_change = db.Column(db.DateTime(timezone=True), server_default=func.now(), onupdate=func.now())

    def __init__(self, username, password):
        self.username = username
        self.password = password

    def __repr__(self):
        return f'<User {self.username}>'

    @staticmethod
    def is_valid_username(username):
        return bool(re.match("^[a-zA-Z]+$", username))