# Ansible

This document detail the steps to follow to initialise a server.

## Pre-Requisites

### on host:
- ssh
- sshpass
- ansible

## Installation

On the server, install ssh:

```bash
sudo apt install openssh-server
```

now, allow the connection with root, in order to to that, modify the line `#PermitRootLogin NotPermited` on the file `/etc/ssh/sshd_config` as followed:

```txt
PermitRootLogin yes
```

if this line is not in the file, just write it.
Now, just reload the ssh service

```bash
service sshd restart
```

Now, setup the root password
```
sudo su
passwd
```

Now, on the host, clone the repository and navigate to the ansible folder

```bash
git clone git@github.com:alexandr-io/backend.git
cd backend/ansible
```

Create the ansible vault secret for the VM's root password

```bash
echo -n '<vm_root_password>' | ansible-vault encrypt_string
```

you should have something like that

```txt
!vault |
    $ANSIBLE_VAULT;1.1;AES256
    545454536473868736578365783657365638756386535375683765387
    345676789898765543456787675643456788987865434567787657...
```

Copy it and replace the value of `ansible_password` in `group_vars/all.yml`.

Finally, fill the `production` file with your servers ip or domains name and run the ansible file with the following command:

```bash
ansible-playbook main.yml -i production --ask-vault-pass
```
