# Clone your fork to your local machine
git clone <fork-url>

# Add 'upstream' repo to list of remotes
git remote add upstream <upstream-url>

# Verify the new remote named 'upstream'
git remote -v

# Fetch from upstream remote
git fetch upstream

# View all branches, including those from upstream
git branch -va

# Checkout your master branch and merge upstream
git checkout master
git merge upstream/master