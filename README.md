# XWS-BSEP
XWS and Security - XML projekat


Za pokretanje je potrebno instalirati **Go** jezik i **Goland** okruzenje za back, a za front(**Angular**) uraditi npm install.

Potrebno je skinuti Nats server i pokrenuti ga pre nego sto se pokrenu mikroservisi sa linka:``` https://github.com/nats-io/nats-server/releases/```

Ukoliko se menjaju proto fajlovi, potrebno je skinuti **ProtoC** i ukucati sledece komande u CMD:
```
- go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
- go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
Moguce je napraviti konfiguraciju kako bi se pokrenuli svi mikroservisi odjednom tako sto se prvo pokrenu pojedinacno i nakon toga se podesi u "_Edit configuration_" u Golandu.

Komanda za ubijanje procesa "_vmmem_" kada se pokrene Docker: ```wsl --shutdown```

Link do originalnog PKI projekta: ``` https://github.com/piwneuh/BSEP-PKI ```

HTTPS je iskljucen medju poslednjim komitima zbog dokera, mogu se videti izmene u komitima u koliko je potrebno da ponovo se ukljuci(postoji mogucnost da je sertifikat istekao).
Sertifikate je potrebno instalirati na svoju lokalnu masinu tako sto se ekstenzija "_.pem_" promeni u "_.cert_" i nakon toga se instalira - vratiti cert u pem nakon toga.