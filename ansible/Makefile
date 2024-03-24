ansible-setup:
	ansible-playbook playbooks/setup.yml -i inventory.yml -vv

ansible-setup-dry:
	ansible-playbook playbooks/setup.yml -i inventory.yml -vvvv --check

ansible-deploy-files:
	ansible-playbook playbooks/deploy_files.yml -i inventory.yml -vv

ansible-firewall:
	ansible-playbook playbooks/firewall.yml -i inventory.yml -vv
