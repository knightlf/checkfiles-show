# checkfiles-show
check something for chongren show 

1. check file and data table from md5.
2. if get exception to update attack data set +1.
3. have timed task from 3 sec.

------------------------
protection resource
sql example:: 

    use crshow;
    
    mysql> update crshow set ca='aaaaaaaaaaaaaaaaa';
    
    Shortcut::
        /data/crshow/checkfiles hsql
    

file example::

    /chongren/test1
    
    /chongren/test2

    Shortcut:
      echo 'aaaa' >> /chongren/test1
    
    
------------------------
run in server:

root@k8s-s3:/data/crshow# ./checkfiles rsql >>/data/logs/crshow/log-$(date +\%Y-\%m-\%d).log &

root@k8s-s3:/data/crshow# ./checkfiles rfile >>/data/logs/crshow/log-$(date +\%Y-\%m-\%d).log &
