_ = require 'underscore'
_.str = require 'underscore.string'
_.mixin(_.str.exports())
Backbone = require 'Backbone'


class App extends Backbone.Router

  routes:
    "": "index"


$ ->
  app = new App()
  Backbone.history.start()
