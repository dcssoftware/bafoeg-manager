# Merge updates from application template into this repo

https://stackoverflow.com/a/56577320

```bash
# add template remote as template-origin
git remote add template-origin git@github.com:dcssoftware/app-template.git

# create new branch chore/rebase-template with current date and switch to it
branch_name="chore/rebase-template-$(date +%Y-%m-%d)"
git branch $branch_name
git checkout $branch_name

# fetch everything from template-origin
git fetch template-origin main

# merge template-origin/main into current branch
git merge template-origin/main --allow-unrelated-histories
```
