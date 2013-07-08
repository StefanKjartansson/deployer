Backbone = require 'Backbone'

exports.Model = class Project extends Backbone.Model

  url: ->
    id = @get 'id'
    "/projects/#{id}/"


exports.Collection = class ProjectCollection extends Backbone.Collection
  model: Project
  url: '/projects/'

exports.CollectionView = class ProjectCollectionView extends Backbone.View

  initialize: (options) ->
    @template = Handlebars.templates.projects
    @listenTo(@model, 'sync', @render)
    @model.fetch()

  render: ->
    @$el.html @template @model.toJSON()
    @


exports.DetailView = class ProjectDetailView extends Backbone.View

  initialize: (options) ->
    @template = Handlebars.templates.project
    @listenTo(@model, 'sync', @reader)
    @model.fetch()

  render: ->
    @$el.html @template @model.toJSON()
    @
