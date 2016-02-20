FROM ehazlett/interlock
COPY ./*.pem /
CMD ["--swarm-url","tcp://172.31.57.87:2376","--swarm-tls-ca-cert","/ca.pem","--swarm-tls-cert","/cert.pem","--swarm-tls-key","/key.pem","--plugin","haproxy","start"]
EXPOSE 80
