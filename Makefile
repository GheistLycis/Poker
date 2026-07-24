PHONY: rebase diff

rebase:
	git checkout main
	git pull
	git checkout -
	git rebase main
	git push --force-with-lease

diff:
	git diff --stat --cached