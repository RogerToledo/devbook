$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar publicação");
    })
}

function curtirPublicacao(evento) {
    evento.preventDefault();
    const eventoCurtir = $(evento.target)
    const publicacaoID = eventoCurtir.closest('div').data('publicacao-id')
    
    $.ajax({
        url: `/publicacoes/curtir/${publicacaoID}`,
        method: "POST",
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao curtir publicação!");
    })
}