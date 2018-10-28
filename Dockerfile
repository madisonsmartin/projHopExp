FROM alpine:3.2
ADD projHopExp-srv /projHopExp-srv
ENTRYPOINT [ "/projHopExp-srv" ]
