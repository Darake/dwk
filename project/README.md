# DBaaS vs DIY

## Cost
The DIY solution seems to be cheaper (at least on some occasions). For example the DBaaS solution for postgres has the cost of $0.090 per GB/month for HDD storage
while with the DIY solution you can get persistant storage as cheap as $0.040 per GB. Though when getting extra features or upgrading to SSD the costs seem to
balance out a bit.

## Required Work
It should come as no suprise that DBaaS solution is easier to get it running. But on can compare the two guides to see this:
https://cloud.google.com/sql/docs/postgres/quickstart  
https://cloud.google.com/architecture/deploying-highly-available-postgresql-with-gke

## Backup
We can see how easy it is to backup or create scheduled backups here: https://cloud.google.com/sql/docs/postgres/backup-recovery/backing-up.  
Its so simple that even a complete novice could do it. In a DIY solution you'd have to do it all by yourself by creating scripts, cronjobs, poking it manually etc.

## My decision
I went with DIY because I already set it up and have no time right now to switch the implementation. If i'd start a real project fresh, I'd propably just go with
DBaaS to save future headache.
