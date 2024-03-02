import os

from flask import Flask, request

app = Flask(__name__)

app.debug = False

WORK_DIR = '/app'
GIT_REPO = 'https://github.com/tasnimzotder/tchat.git'


def read_and_strip(command):
    return command.read().split('\n')[0].strip()


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

        # create a file acme.json for traefik if not exists
        if not os.path.exists('/app/tchat/acme.json'):
            with open('/app/tchat/acme.json', 'w') as f:
                f.write('{}')

        # build and run the docker containers
        os.system(
            "docker compose -f {}/tchat/compose.yaml up -d --build".format(WORK_DIR))

        res = {
            "status": "OK",
            "message": "Deployment triggered"
        }

        return res, 200

    else:
        res = {
            "status": "ERROR",
            "message": "Invalid webhook"
        }

        return res, 400


@app.route('/health', methods=['GET'])
def health():
    os_name = os.popen('uname').read().split('\n')[0]

    system_info = {}
    container_info = {}

    if os_name == 'Darwin':
        # get system information
        system_info["uptime"] = read_and_strip(os.popen('uptime'))
        system_info["memory_usage"] = read_and_strip(
            os.popen('top -l 1 | grep PhysMem'))
        system_info["cpu_usage"] = read_and_strip(
            os.popen('top -l 1 | grep CPU'))

    elif os_name == 'Linux':
        system_info["uptime"] = read_and_strip(os.popen('uptime -p'))
        system_info["memory_usage"] = read_and_strip(os.popen('free -m'))
        system_info["cpu_usage"] = read_and_strip(
            os.popen('top -bn1 | grep load'))

    # get container information
    container_info["container_count"] = read_and_strip(
        os.popen('docker ps -aq | wc -l'))
    container_info["running_container_count"] = read_and_strip(
        os.popen('docker ps -q | wc -l'))
    container_info["image_count"] = read_and_strip(
        os.popen('docker images -q | wc -l'))
    container_info["volume_count"] = read_and_strip(
        os.popen('docker volume ls -q | wc -l'))
    container_info["network_count"] = read_and_strip(
        os.popen('docker network ls -q | wc -l'))

    res = {
        "status": "OK",
        "version": "0.0.1",
        "system_info": system_info,
        "container_info": container_info,
    }

    return res, 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000)
