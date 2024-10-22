# Borzoi

### Overview

**Borzoi** is a lightweight deployment tool designed to simplify the process of deploying applications to your private server (for those who usually do it by hand).

### How to Use Borzoi

1. In the root of your repository, create a `borzoi` directory and place a `borzoi.yaml` file inside. Add a `docker-compose.yaml` to manage your application’s container setup.
2. Ensure both `nginx` and `docker` are installed and configured on your server.
3. Run Borzoi with `sudo` on your server (e.g., `borzoi.server.com`).
4. Send a POST request to `borzoi.server.com/api/v1/deploy/create`, including your repository URL and branch name. You can also set up GitHub Webhooks to automate this step.
5. Visit `borzoi.server.com` to see your application running.

### License

This project is licensed under the MIT License.  
You can view the full license in the [LICENSE](./LICENSE) file.
