function links(parent, args, context) {
    return context.prisma.user({ id: parent.id }).links()
}

function votes(parent, args, context) {
    return context.prisma.user({ id: parent.id }).votes()
}
  
module.exports = {
    id: (parent) => parent.id,
    name: (parent) => parent.name,
    email: (parent) => parent.email,
    links,
    votes,
}