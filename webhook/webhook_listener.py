import os

from flask import Flask, request

app = Flask(__name__)

app.debug = False

WORK_DIR = '/app'
GIT_REPO = 'https://github.com/tasnimzotder/tchat.git'


@app.route('/webhook', methods=['POST'])
def handle_webhook():
    if request.json and request.json.get('ref') == 'refs/heads/main':
        os.system("cd {}".format(WORK_DIR))

        # clone the repository "tchat" if not exists
        if not os.path.exists('/app/tchat'):
            os.system("git clone {}".format(GIT_REPO))

        # pull the latest changes from the repository
        os.system("cd {}/tchat".format(WORK_DIR))
        os.system("git pull origin main")

        # build and run the docker containers
        os.system("docker compose -f {}/tchat/server/compose.yaml up -d --build".format(WORK_DIR))

        return 'Deployment triggered', 200
    else:
        return 'Invalid webhook', 400


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000)
