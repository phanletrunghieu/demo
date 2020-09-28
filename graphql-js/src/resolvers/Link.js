function postedBy(parent, args, context) {
    return context.prisma.link({ id: parent.id }).postedBy()
}

function votes(parent, args, context) {
    return context.prisma.link({ id: parent.id }).votes()
}

module.exports = {
    id: (parent) => parent.id,
    description: (parent) => parent.description,
    url: (parent) => parent.url,
    postedBy,
    votes,
}