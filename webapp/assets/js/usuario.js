$('#parar-seguir').on('click', pararSeguir)
$('#seguir').on('click', seguir)

function pararSeguir() {
    const usuarioID = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/parar-seguir/${usuarioID}`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioID}`;
    }).fail(function() {
        Swal.fire({
            title: 'Ops...',
            text: 'Erro ao parar de seguir usuário!',
            icon: 'error'
        });
        $('#parar-seguir').prop('disabled', false);
    })
}

function seguir() {
    const usuarioID = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/seguir/${usuarioID}`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioID}`;
    }).fail(function() {
        Swal.fire({
            title: 'Ops...',
            text: 'Erro ao parar de seguir usuário!',
            icon: 'error'
        });
        $('#seguir').prop('disabled', false);
    })
}