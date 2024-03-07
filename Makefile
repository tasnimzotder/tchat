act:
	act --container-architecture linux/amd64 --secret-file .env.local

act-ci:
	act --container-architecture linux/amd64 \
	-W ./.github/workflows/ci.yml

act-cd:
	act --container-architecture linux/amd64 \
	--secret-file .env.local -W ./.github/workflows/cd.yml
