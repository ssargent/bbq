from flask import Flask,jsonify
import os
app = Flask(__name__)

@app.route("/")
def hello():
    return jsonify(os.environ)
    
if __name__ == "__main__":
    app.run()