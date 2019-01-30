$(document).ready(function() {
    console.log("Bienvenido a Instagram");
    var formularioRegistro = $("#formularioRegistro > form > input:last-child");
    console.log(formularioRegistro);

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
            $('#formularioRegistro > p').html('<strong>Su petición ha sido registrada con éxito</strong>');
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });


    // Esto es el registro para que cuando pulses un boton se te muestre un formulario u otro

    $('#formularioConectate').hide();
    $('#formularioRegistro').hide();


    $('.registrate').on ('click', function () {
        console.log('click en el boton regisro');
        $('#formularioConectate').hide();
        $('#formularioRegistro').show();
        $('#botones').hide();
    })

    $('.conectate').on ('click', function () {
        console.log('click en el boton regisro');
        $('#formularioConectate').show();
        $('#formularioRegistro').hide();
        $('#botones').hide();
    })




});