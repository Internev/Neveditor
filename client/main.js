import Vue from 'vue'
import Resource from 'vue-resource'

var channel, conn;

Vue.use(Resource)

new Vue({
  el: '#container',

  data: {
    ws: null,
    tesisdoc: 'doccotexteorooni',
    channel: '',
  },

  created: function() {
    var self = this;
    this.$http.get('/getUrl').then((res) => {
      // this is to get the unique url for the websocket. Need to assign connection inside of this call because async.
      console.log('getUrl datas!:', res);
      this.channel = res.body;
      this.ws = new WebSocket('ws://' + window.location.host + '/ws/' + this.channel);
      // var sessionId = null;

      console.log('websocket address is:',this.ws);
      // Textarea is editable only when socket is opened.
      // this.ws.onopen = function(e) {
      //   content.attr("disabled", false);
      // };
      //
      // this.ws.onclose = function(e) {
      //   content.attr("disabled", true);
      // };
      console.log('document contents:', this.tesisdoc);
      // Whenever we receive a message, update textarea
      this.ws.onmessage = (e) => {
        if (e.data != this.tesisdoc) {
          this.tesisdoc = e.data;
        }
      };
    });
  },

  methods: {
    send: function() {
      console.log('sending:', this.tesisdoc);
      this.ws.send(this.tesisdoc)
    }
  }
})
