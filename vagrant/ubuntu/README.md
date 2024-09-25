# VAGRANT COM ANSIBLE + WORDPRESS

## Passo a Passo

### VAGRANT

O presente vagrant file utiliza-se de uma imagem ubuntu de versão "Ubuntu 18.04.6 LTS (GNU/Linux 4.15.0-212-generic x86_64)" com ip fixo e público.

### ANSIBLE

A ideia de utilizar o ansible nesse projeto foi a de automatizar a instalação de diversas dependências e o Wordpress na concepção da vm. Porém por enquanto ele faz apenas isso, espera-se em versões futuras que ele possa configurar o wordpress também.

### WORDPRESS

Após a inicialização da VM, há a necessidade de entrar e configurar o wordpress assim como seria manualmente.

- 1. Configure Apache para WordPress

Crie um site Apache para o wordpress no seguinte destino: `etc/apache2/sites-available/wordpress.conf` com o código:

```HTML

<VirtualHost *:80>
    DocumentRoot /srv/www/wordpress
    <Directory /srv/www/wordpress>
        Options FollowSymLinks
        AllowOverride Limit Options FileInfo
        DirectoryIndex index.php
        Require all granted
    </Directory>
    <Directory /srv/www/wordpress/wp-content>
        Options FollowSymLinks
        Require all granted
    </Directory>
</VirtualHost>


```

---

Habilite o site com:

`sudo a2ensite wordpress`

Habilite a reescrita de URL com:

`sudo a2enmod rewrite`

Desabilite o site padrão "It Works" com:

`sudo a2dissite 000-default`

Reinicie o serviço para aplicar as mudanças:

`sudo service apache2 reload`

---

- 2. Configurandoo banco de dados:

Para configurar o wordpress precisamos criar um BANCO DE DADOS MySQL:

> Certifique-se de alterar o campo <your-password>

```bash

$ sudo mysql -u root
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 7
Server version: 5.7.20-0ubuntu0.16.04.1 (Ubuntu)

Copyright (c) 2000, 2017, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> CREATE DATABASE wordpress;
Query OK, 1 row affected (0,00 sec)

mysql> CREATE USER wordpress@localhost IDENTIFIED BY '<your-password>';
Query OK, 1 row affected (0,00 sec)

mysql> GRANT SELECT,INSERT,UPDATE,DELETE,CREATE,DROP,ALTER
    -> ON wordpress.*
    -> TO wordpress@localhost;
Query OK, 1 row affected (0,00 sec)

mysql> FLUSH PRIVILEGES;
Query OK, 1 row affected (0,00 sec)

mysql> quit
Bye

```

> Tenha em mente que este é apenas um exemplo e que é recomendado que os valores sejam alterados

habilite então o MySQL:

`sudo service mysql start`

---

- 3. Configurando a conexão do Banco de dados para o Wordpress

Agora é preciso configurar o Wordpress para usar este banco de dados, primeiro precisamos garantir que os arquivos e pastas que vamos mexer tenham a devida permissão.

```bash

sudo chown -R www-data:www-data /srv/www/wordpress
sudo chmod -R 755 /srv/www/wordpress


```

Agora vamos copiar o template de configuração:

`sudo -u www-data cp /srv/www/wordpress/wp-config-sample.php /srv/www/wordpress/wp-config.php`

Em seguida, defina as credenciais do banco de dados no arquivo de configuração (não substitua database_name_here ou username_here nos comandos abaixo. Substitua <your-password> pela senha do seu banco de dados.):

```bash

sudo -u www-data sed -i 's/database_name_here/wordpress/' /srv/www/wordpress/wp-config.php
sudo -u www-data sed -i 's/username_here/wordpress/' /srv/www/wordpress/wp-config.php
sudo -u www-data sed -i 's/password_here/<your-password>/' /srv/www/wordpress/wp-config.php

```

Por fim, em uma sessão de terminal, abra o arquivo de configuração no vim:

`sudo -u www-data vim /srv/www/wordpress/wp-config.php`

Aqui você precisará mudar alguns parâmetros assim como eu fiz. No campo de senha logo no início do arquivo, troque o campo "<your-password>" pela senha que você definiu.

Em seguida procure as linhas:

```bash

define( 'AUTH_KEY',         'put your unique phrase here' );
define( 'SECURE_AUTH_KEY',  'put your unique phrase here' );
define( 'LOGGED_IN_KEY',    'put your unique phrase here' );
define( 'NONCE_KEY',        'put your unique phrase here' );
define( 'AUTH_SALT',        'put your unique phrase here' );
define( 'SECURE_AUTH_SALT', 'put your unique phrase here' );
define( 'LOGGED_IN_SALT',   'put your unique phrase here' );
define( 'NONCE_SALT',       'put your unique phrase here' );

```

e as mude de acordo com o [link](https://api.wordpress.org/secret-key/1.1/salt/)

---

- 4. Pronto, agora você pode configurar seu wordpress acessando o ip público da VM.

---

> Todas essas informações foram tiradas, traduzidas e reduzidas de acordo com a minha necessidade do site: [ubuntu.com](https://ubuntu.com/tutorials/install-and-configure-wordpress#7-configure-wordpress)
