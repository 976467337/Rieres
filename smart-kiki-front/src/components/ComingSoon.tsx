interface ComingSoonProps {
  title: string
}

export function ComingSoon({ title }: ComingSoonProps) {
  return (
    <div className="flex flex-1 flex-col items-center justify-center gap-2 px-6 py-16 text-center">
      <h2 className="font-heading text-lg font-bold">{title}</h2>
      <p className="text-sm text-muted-foreground">Essa área ainda está em construção. Em breve por aqui!</p>
    </div>
  )
}
