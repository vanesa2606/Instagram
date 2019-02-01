$(document).ready(function() {
    console.log("Bienvenido a Instagram");
    var formularioRegistro = $("#formularioRegistro > input:last-child");
    console.log(formularioRegistro);

    //La linea de abajo es para saber si me ha iniciado sesión y me ha creado una cookie
    //console.log(document.cookie);

    // Ajax para el regitro
    $(formularioRegistro).click(function() {
        var user = $("#nombre").val();
        var usernam = $("#username").val();
        var electronico = $("#correo").val();
        var password = $("#contrasena").val();
        console.log(nombre, username, correo, contrasena);
        
        var envio = {
            nombre: user,
            username: usernam,
            correo: electronico,
            contrasena: password
        };
        console.log(envio);
        $.post({
            url:"/registro",
            data: JSON.stringify(envio),
            method: "POST",
            success: function(data, status, jqXHR) {
                console.log(data);
            },      
            dataType: "json"


        }).done(function(data) {
            console.log("Petición realizada");
            $('#index > #botones > h1').html('Su petición ha sido registrada con éxito!!!');

        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });


    // Ajax para el login
    var formularioConectate = $("#formularioConectate > input:last-child");
    console.log(formularioConectate)
    $(formularioConectate).click(function() {
        var usernam = $("#usernamecone").val();
        var password = $("#contrasenacone").val();
        console.log(username, contrasena);
        
        var envio = {
            username: usernam,
            contrasena: password
        };
        console.log(envio);
        $.post({
            url:"/login",
            data: JSON.stringify(envio),
            method: "POST",
            success: function(data, status, jqXHR) {
                console.log(data);
            },      
            dataType: "json"


        }).done(function(data) {
            console.log("Petición realizada");
            if (data == true){
                window.location.href = "/principal";
            }
        
        }).fail(function(data) {
            console.log("Petición fallida");
            console.log(data);

        
        }).always(function(data){
            console.log("Petición completa");
        });
    });



    // Esto es el registro para que cuando pulses un boton se te muestre un formulario u otro

    $('#formularioConectate').hide();
    $('#formularioRegistro').hide();


    $('.registrate').on ('click', function () {
        $('#formularioConectate').hide();
        $('#formularioRegistro').show();
        $('#botones').hide();
    })

    $('.conectate').on ('click', function () {
        $('#formularioConectate').show();
        $('#formularioRegistro').hide();
        $('#botones').hide();
    })




});