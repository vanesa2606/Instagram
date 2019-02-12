$(document).ready(function() {
    console.log("Bienvenido a Instagram");
    var formularioRegistro = $("#formularioRegistro > #registro");
    console.log(formularioRegistro);

    //La linea de abajo es para saber si me ha iniciado sesión y me ha creado una cookie
    console.log(document.cookie);

    // LLamo a la funcion de listar fotos
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
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });


    // Ajax para el login
    var formularioConectate = $("#formularioConectate > input:last-child, #formularioRegistro > button");
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


    // Añadir los comentarios 
    $('#publicaciones').on('click', '.btnEnviarCom', function() {
        console.log(this);
        var comentario = $(this).parent().find(".txtComentario").val();
        var id = $(this).siblings(".txtId").val();
        console.log(comentario, id);
        var envio = {
            texto: comentario,
            id: id
        };
        $.post({
            url:"/comentario",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#txtTexto").val('')
                $('#txtId').val('')
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });

    });
    
        
       
});

 //Imprimir todas las fotos por pantalla 

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
                        "<input type='hidden' class='txtId' value="+array[x].ID+">"+
                        "<img src='/files/"+array[x].URL+"' width='95%'>"+
                        "<div class='container'>"+
                            "<h3>"+ array[x].Texto + "</h3>"+
                        "</div>"+
                        "<div class='comentarios'>"+
                        "</div>"+
                        "<input type='comentario' class='txtComentario'>"+
                        "<input type='button' class='btnEnviarCom' value='Enviar'>"+
                    "</div>");
                    // Llamo a la función de mostrar comentarios

                    MostrarComentarios();

            }
        } else {
            tbody.append('<tr><td colspan="3">No hay publicaciones que mostar</td></tr>');
            
        }
}


 // Imprimir todas los comentarios por pantalla 

 function MostrarComentarios() {
    $.ajax({
        url: "/listarcomentario",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
        Comentario(data); 
        console.log(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Comentario(array) {
    var div = $("#publicaciones .comentarios");
    div.children().remove();
        if(array != null && array.length > 0) {

            for(var x = 0; x < array.length; x++) {
                div.append(
                    "<input type='hidden' class='txtId' value="+array[x].ID+">"+
                    "<div class='container'>"+
                        "<h6>"+ array[x].Username + "</h6>"+
                        "<p>"+ array[x].Texto + "</p>"+
                        "<hr align='center' noshade='noshade' size='1' width='80%' />"+
                    "</div>");
            }
        } else {
            tbody.append('<tr><td colspan="3">No hay publicaciones que mostar</td></tr>');
            
        }
}



