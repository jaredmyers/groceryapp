import json
from flask import Flask, jsonify
import mysql.connector

app = Flask(__name__)
@app.route("/")


def index():
    return jsonify ({'name': 'alice', 'email': 'alice@outlook.com'})

app.run()
