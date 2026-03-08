import { createFileRoute } from '@tanstack/react-router'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'

export const Route = createFileRoute('/(private)/companies')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="bg-linear-to-br from-[#FFF3E6] to-[#F3E8FF] h-[calc(100vh-64px)] p-8 gap-4 flex flex-col">
      <div className="max-w-2xl mx-auto w-full">
        <Card className="border-2 border-[#E5E7EB] shadow-lg">
          <CardHeader>
            <CardTitle className="text-xl font-bold text-[#1F2937]">Crear Nueva Compañía</CardTitle>
          </CardHeader>
          <CardContent>
            <form className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="name">Nombre de la compañía</Label>
                <Input
                  id="name"
                  placeholder="Ej: TechCorp"
                  className="border-[#E5E7EB] focus:border-[#F97316]"
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="description">Descripción</Label>
                <Input
                  id="description"
                  placeholder="Breve descripción de la empresa"
                  className="border-[#E5E7EB] focus:border-[#F97316]"
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="industry">Industria</Label>
                <Input
                  id="industry"
                  placeholder="Ej: Tecnología, Salud, Educación..."
                  className="border-[#E5E7EB] focus:border-[#F97316]"
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="website">Sitio web</Label>
                <Input
                  id="website"
                  placeholder="https://ejemplo.com"
                  type="url"
                  className="border-[#E5E7EB] focus:border-[#F97316]"
                />
              </div>
              
              <div className="space-y-2">
                <Label htmlFor="location">Ubicación</Label>
                <Input
                  id="location"
                  placeholder="Ciudad, País"
                  className="border-[#E5E7EB] focus:border-[#F97316]"
                />
              </div>
              
              <Button 
                type="submit" 
                className="w-full bg-[#F97316] hover:bg-[#EA580C] text-white font-semibold"
              >
                Crear Compañía
              </Button>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  )
}
