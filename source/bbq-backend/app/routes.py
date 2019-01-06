from flask import render_template
from app import app

@app.route('/')
@app.route('/index')
def index():
    user = {'username': 'scott'}
    cooks = [
        {
            'components': {'main': 'Boston Butt'},
            'description': 'Slow smoked boston butt'
        },
        {
            'components': {'main': 'Whole Chicken'},
            'description': 'Slow smoked whole chicken'
        }
    ]
    return render_template('index.html', title='Home', user=user, cooks=cooks)  