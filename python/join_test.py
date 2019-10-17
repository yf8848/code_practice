
PRODUCT = 'aweme'
SUBSYS = 'poi'
MODULE = 'cron_data_differ'

psm = '.'.join([PRODUCT, SUBSYS, MODULE])


print(isinstance(psm, str), psm)