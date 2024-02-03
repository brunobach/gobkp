GOBKP - CLI para Backup em Golang
=================================

GOBKP é uma CLI escrita em Go (Golang) para facilitar o backup de arquivos e diretórios.

Funcionalidades
---------------

*   **Backup Simples**: Faça backup de arquivos e diretórios com facilidade.
*   **Restauração Simples**: Restaure arquivos e diretórios de um arquivo zip de backup.
*   **Configuração Flexível**: Use um arquivo de configuração para especificar os arquivos a serem incluídos e os padrões de exclusão.

Instalação
----------

1.  **Instale Go**: Certifique-se de ter o Go instalado em sua máquina. Você pode encontrar instruções de instalação em golang.org.
    
2.  **Instale a CLI**: Clone este repositório ou execute o seguinte comando para instalar a CLI:
    
```bash
go install github.com/brunobach/gobkp/cmd/gobkp@latest
```

Uso
---

### Backup


```bash
gobkp backup
```

Isso criará um arquivo zip de backup com base nas configurações especificadas no arquivo `backup.cfg`.

### Restauração

```bash
gobkp restore
```

Isso restaurará os arquivos e diretórios do arquivo zip de backup.

Configuração
------------

Você pode configurar os arquivos a serem incluídos e os padrões de exclusão no arquivo `backup.cfg`.

Exemplo de `backup.cfg`:

plaintextCopy code

```plain
# Arquivos a serem incluídos no backup 
~/.zshrc
~/.zsh_history
~/.ssh
~/.gitconfig
~/.vault-token
~/.aws 

# Padrões de exclusão (ignorar esses arquivos durante o backup) 
exclude: 
~/ssh/id_rsa 
~/.aws/credentials.bak 
```
Contribuindo
------------

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

* * *
