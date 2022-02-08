var app = {
    settings: {
      container: $('.calendar'),
      calendar: $('.front'),
      days: $('.weeks span'),
      form: $('.back'),
      input: $('.back input'),
      buttons: $('.back button')
    },
  
    init: function() {
      instance = this;
      settings = this.settings;
      this.bindUIActions();
    },
  
    swap: function(currentSide, desiredSide) {
      settings.container.toggleClass('flip');
  
      currentSide.fadeOut(900);
      currentSide.hide();
      desiredSide.show();
  
    },
  
    bindUIActions: function() {
      settings.days.on('click', function(){
        instance.swap(settings.calendar, settings.form);
        settings.input.focus();
      });
  
      settings.buttons.on('click', function(){
        instance.swap(settings.form, settings.calendar);
      });
    }
  }
  
  app.init();