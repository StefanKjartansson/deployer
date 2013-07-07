Backbone = require 'Backbone'

exports.Model = class Project extends Backbone.Model

exports.Collection = class ProjectCollection extends Backbone.Collection
  model: Project
  url: '/projects/'

exports.CollectionView = class ProjectCollectionView extends Backbone.View

  initialize: (options) ->
    @template = Handlebars.templates.eventContainer
    @listenTo(@model, 'change', @render)
    @model.fetch()

  render: ->
    @$el.html @template {}

