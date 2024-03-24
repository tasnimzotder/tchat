import os

from flask import Flask, request, jsonify
from waitress import serve
import logging

app = Flask(__name__)
app.debug = False

WORK_DIR = '/app'
GIT_REPO = 'https://github.com/tasnimzotder/tchat.git'


def strip_str(string: str) -> str:
    return string.split('\n')[0].strip()


def read_and_strip(command) -> str:
    """
    Function to read and strip the input command and return the first line after splitting by newline and stripping any leading/trailing whitespaces.
    """
    command_str = command.read()

    return strip_str(command_str)


@app.before_request
def log_request() -> None:
    app.logger.info("%s %s %s", request.method,
                    request.path, request.remote_addr)


@app.after_request
def log_response(response) -> None:
    if response.status_code != 200:
        app.logger.info('%s', response.status)

    return response


@app.route('/webhook', methods=['POST'])
def handle_webhook() -> None:
    """
    Handle the webhook request and trigger the deployment process if the webhook is valid. Returns a status message and HTTP status code.
    """
    if request.json and request.json.get('ref') == 'refs/heads/main':
        file_name = "deploy.sh"

        file_path = "{}/{}".format(WORK_DIR, file_name)
        print(file_path)

        # check if the file exists
        if not os.path.exists(file_path):
            res = {
                "status": "ERROR",
                "message": "File not found"
            }
            return jsonify(res), 404

        # read the exit code
        exit_code = os.system("bash {}".format(file_path))
        if exit_code != 0:
            res = {
                "status": "ERROR",
                "message": "Deployment failed"
            }
            return res, 500

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

        return jsonify(res), 400


@app.route('/health', methods=['GET'])
def health() -> None:
    """
    This function retrieves health information about the system and containers and returns the information along with a status code.
    """
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

    return jsonify(res), 200


@app.route('/', methods=['GET'])
def index() -> None:
    res = {
        "status": "OK",
        "message": "Webhook Listener"
    }

    return jsonify(res), 200


def main() -> None:
    app.logger.setLevel(logging.INFO)
    serve(app, host='0.0.0.0', port=3000)


if __name__ == '__main__':
    main()
