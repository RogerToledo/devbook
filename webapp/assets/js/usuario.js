$('#parar-seguir').on('click', pararSeguir)

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

}