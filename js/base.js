$(document).ready(function() {
    console.log("Bienvenido a Instagram");
    var formularioRegistro = $("#formularioRegistro > input:last-child");
    console.log(formularioRegistro);

    //La linea de abajo es para saber si me ha iniciado sesión y me ha creado una cookie
    console.log(document.cookie);
    MostrarFotosssss();
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
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });
 

    // Poner en el boton de cerrar sesión la funcion de logout
    
    $('#cerrarSesion').click(function() {
        window.location.href = "/logout";
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

    // Comentarios
    var botonComen = $("#publicaciones .btnEnviarCom");
    console.log("boton de envio", botonComen);
    $(botonComen).click(function() {
        var comentario = $(this).parent().children().children().find(".txtComentario").val();
        console.log("Has pulsado el botón del envio de comentario", comentario);
        
       /* var envio = {
            palabra: texto
        };

        $.post({
            url:"/",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#txtTexto").val('')
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            ActualizarHistorial();
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });*/
    });

});

 // Imprimir todas las fotos por pantalla 

 function MostrarFotosssss() {
    $.ajax({
        url: "/listarfoto",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
        Fotos(data); 
        console.log(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Fotos(array) {
var tbody = $("#publicaciones");
tbody.children().remove();
    if(array != null && array.length > 0) {

        for(var x = 0; x < array.length; x++) {
            tbody.append(
                "<div class='card'>"+
                    "<img src='/files/"+array[x].URL+"' width='15%'>"+
                    "<div class='container'>"+
                        "<p>"+ array[x].Texto + "</p>"+
                    "</div>"+
                        "<input type='comentario' class='txtComentario'>",
                        "<input type='button' class='btnEnviarCom' value='Enviar'>",
                "</div>");
        }
    } else {
        tbody.append('<tr><td colspan="3">No hay registros de hoy</td></tr>');
        
    }
}


