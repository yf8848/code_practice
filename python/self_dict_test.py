# -*- coding: utf-8 -*-
import time
import uuid
import json


class BaseMessage(object):
    topic = ''
    conf = "nsq.NSQ_CONF3"
    cluster = "nsq.NSQ_CLUSTER_LF_TEST"

    def __init__(self, **kws):
        self.event_time = kws.get('event_time', int(time.time()))
        self.msg_id = kws.get('msg_id', str(uuid.uuid1()).replace('-', ''))
        self.addr = kws.get('addr', "get_local_ip")
        self.caller = kws.get('caller', "get_psm")
        self.log_id = kws.get('log_id', "get_current_logid")
        self.is_perf_test = kws.get('is_perf_test', False)
        self.origin_dc = kws.get('origin_dc', "UNKNOWN")

    def to_dict(self):
        return self.__dict__


    def send_to_mq(self, defer=0):
        producer = None
        msg = "None"
        try:
            #producer = NsqProducer(self.topic, cluster=self.cluster, stress_testing_strategy=stress_testing_strategy)
            msg = json.dumps(self.to_dict())
            print("%s" % msg)
        except Exception as e:
            return False

    def __str__(self):
        return '%s' % self.to_dict()

    def __repr__(self):
        return self.__str__()

class MusicianMusicPostActionMessage(BaseMessage):
    topic = "musician_music_post_action"
    conf = "nsq.NSQ_IES1"
    cluster = "nsq.NSQ_CLUSTER_LF_IES_NORMAL"

    def __init__(self, **kwargs):
        super(MusicianMusicPostActionMessage, self).__init__(**kwargs)
        self.process_mode = kwargs.get('process_mode')
        self.song_ids = kwargs.get('song_ids')
        self.operator = kwargs.get('operator')
        self.remark = kwargs.get('remark')

class DeauthenticationV2Handler(object):
    
    def __init__(self):
        self.request = "request"
        self.user_id = "request.UserId"
        self.process_mode = "request.MusicianMusicProcessMode"
        self.operator = "request.Operator"
        self.user = None
    def send_test(self):
        MusicianMusicPostActionMessage(process_mode=self.process_mode, song_ids="process_song_ids",
                                               operator=self.operator, remark=u"音乐人注销").send_to_mq()

if __name__ == "__main__":
    DeauthenticationV2Handler().send_test()
    print("%s" % MusicianMusicPostActionMessage().topic)

    print("%s" % MusicianMusicPostActionMessage().__dict__)