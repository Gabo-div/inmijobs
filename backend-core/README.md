si van a hacer tests, aqui les paso los formatos json:
para crear post y editar:
post normal:
{
    "user_id": "user_test_123",
    "title": "Mi fin de semana en el evento tech",
    "content": "Increíble experiencia compartiendo con la comunidad. ¡Aquí les dejo unas fotos!",
    "images": [
        "https://cdn.tusitio.com/fotos/evento_1.jpg",
        "https://cdn.tusitio.com/fotos/evento_2.jpg"
    ],
    "job_id": null,
    "company_id": null
}

post para vacantes:
{
    "user_id": "user_test_123",
    "title": "¡Buscamos Senior Go Developer!",
    "content": "Únete a nuestro equipo de infraestructura. Buscamos expertos en concurrencia y microservicios. Aplicaciones abiertas.",
    "job_id": 1,
    "images": [
        "https://cdn.tusitio.com/banners/hiring_go.png"
    ]
}
"O" 
{
    "user_id": "user_test_123",
    "title": "¡Buscamos Senior Go Developer!",
    "content": "Únete a nuestro equipo de infraestructura. Buscamos expertos en concurrencia y microservicios. Aplicaciones abiertas.",
    "job_id": 1,
    "company_id":1,
    "images": [
        "https://cdn.tusitio.com/banners/hiring_go.png"
    ]
}

(de ambas maneras son validas porque el job_id se vincula automaticamente con el company_id si este no se coloca) 

para crear comentario o editar:
{
    "message": "¡Hola! Gabo, como estas",
    "post_id": "1",
    "user_id": "user_test_123"
}

reaction:

{
	"user_id": "user_test_123",
	"reaction_id":  1 
}