import React, { useState } from "react";
/* eslint-disable sort-imports */
import { 
    Cake, Globe, GraduationCap, Hand, Heart, Home, 
    Lock, Mail, MapPin, Mic, Music, Pencil, Phone, Pin, Plane, Plus, 
    Briefcase, Link as LinkIcon, MessageSquare, User as UserIcon, Users
} from "lucide-react";

import { Button } from "@/components/ui/button";

// --- COMPONENTE BASE PARA CADA FILA (Reutilizable) ---
interface InfoRowProps {
    icon: React.ElementType;
    label?: string;
    value: React.ReactNode;
    subValue?: string;
    action?: boolean; // Si es true, se ve como un botón de "Agregar..."
    rightElement?: React.ReactNode; // Para los iconos de candado/lápiz
}

function InfoRow({ icon: Icon, label, value, subValue, action, rightElement }: InfoRowProps) {
    return (
        <div className={`flex items-start gap-4 p-2 -ml-2 rounded-lg transition-colors ${action ? 'cursor-pointer hover:bg-muted/50 group' : ''}`}>
            {/* Icono Principal */}
            <div className="w-9 h-9 mt-0.5 rounded-full bg-secondary/50 flex items-center justify-center shrink-0 text-muted-foreground">
                <Icon className="w-5 h-5" />
            </div>

            {/* Contenido Central */}
            <div className="flex-1 min-w-0 flex flex-col justify-center min-h-[2.25rem]">
                {label && <p className="text-sm font-semibold text-foreground mb-0.5">{label}</p>}
                
                <div className={`${action ? 'text-blue-600 group-hover:underline font-medium' : 'text-foreground text-sm'}`}>
                    {value}
                </div>
                
                {subValue && <p className="text-xs text-muted-foreground mt-0.5">{subValue}</p>}
            </div>

            {/* Elementos a la derecha (Candado, Lápiz) */}
            {rightElement && (
                <div className="flex items-center gap-2 text-muted-foreground">
                    {rightElement}
                </div>
            )}
        </div>
    );
}

// helper hook for editable text fields
function useField(initial: string, label: string, placeholder?: string, inputType: string = 'text') {
    const [value, setValue] = useState(initial);
    const [open, setOpen] = useState(false);
    const openEditor = () => setOpen(true);
    const closeEditor = () => setOpen(false);
    const display = value || <span className="text-muted-foreground">{placeholder ?? label}</span>;
    const modal = (
        <EditModal
            open={open}
            label={label}
            initialValue={value}
            placeholder={placeholder}
            inputType={inputType}
            onClose={closeEditor}
            onSave={(val) => setValue(val)}
        />
    );
    return { value, display, openEditor, modal, setValue };
}

// --- HELPER PARA ICONOS DE EDICIÓN (Candado/Lápiz) ---
function EditActions({ onEdit }: { onEdit?: () => void }) {
    return (
        <>
            <Button variant="ghost" size="icon" className="h-8 w-8 text-muted-foreground/60 hover:text-foreground">
                <Lock className="w-4 h-4" />
            </Button>
            <Button variant="ghost" size="icon" className="h-8 w-8 text-muted-foreground/60 hover:text-foreground" onClick={onEdit}>
                <Pencil className="w-4 h-4" />
            </Button>
        </>
    );
}

// --- MODAL SIMPLE PARA EDITAR CAMPOS ---
function EditModal({
    open,
    label,
    initialValue,
    placeholder,
    inputType = 'text',
    onClose,
    onSave,
}: {
    open: boolean;
    label: string;
    initialValue?: string;
    placeholder?: string;
    inputType?: string;
    onClose: () => void;
    onSave: (val: string) => void;
}) {
    const [value, setValue] = useState(initialValue ?? "");

    // when opening with a different initial value, sync it
    React.useEffect(() => {
        if (open) setValue(initialValue ?? "");
    }, [open, initialValue]);

    if (!open) return null;

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40">
            <div className="bg-white dark:bg-card rounded-lg p-4 w-full max-w-sm shadow-lg">
                <h3 className="text-lg font-semibold mb-2">Editar {label}</h3>
                <input
                    autoFocus
                    type={inputType}
                    value={value}
                    onChange={(e) => setValue(e.target.value)}
                    placeholder={placeholder}
                    className="w-full px-3 py-2 rounded-md border border-gray-200 dark:border-border bg-transparent mb-4"
                />
                <div className="flex justify-end gap-2">
                    <Button variant="ghost" onClick={onClose}>Cancelar</Button>
                    <Button onClick={() => { onSave(value); onClose(); }}>Guardar</Button>
                </div>
            </div>
        </div>
    );
}

// ================= SECCIONES ESPECÍFICAS =================

export function DetailsSection() {
    const presentation = useField("Información sobre ti", "Presentación");
    const pinned = useField("Detalles fijados", "Detalles fijados");
    return (
        <div className="space-y-6 animate-in fade-in">
            <div>
                <h3 className="text-lg font-bold mb-4">Presentación</h3>
                <InfoRow
                    icon={Hand}
                    value={presentation.value}
                    rightElement={<EditActions onEdit={presentation.openEditor} />}
                />
                {presentation.modal}
            </div>
            <div>
                <h3 className="text-lg font-bold mb-4">Detalles fijados</h3>
                <InfoRow
                    icon={Pin}
                    value={pinned.display}
                    rightElement={<EditActions onEdit={pinned.openEditor} />}
                />
                {pinned.modal}
            </div>
        </div>
    );
}

export function PersonalDataSection() {
    const locationField = useField("Ciudad o localidad actual", "Ubicación");
    const originField = useField("Ciudad de origen", "Ciudad de origen");
    const birthday = useField("", "Fecha de nacimiento", "DD/MM/AAAA", "date");
    const relationship = useField("Situación sentimental", "Situación sentimental");
    const family = useField("Familia", "Familiares");
    const gender = useField("Hombre", "Género");
    const pronouns = useField("masculino", "Pronombres", "masculino / femenino / etc.");
    const languages = useField("Idiomas", "Idiomas");

    return (
        <div className="space-y-6 animate-in fade-in">
            {/* Ubicación */}
            <div>
                <h3 className="font-semibold mb-3">Ubicación</h3>
                <InfoRow icon={MapPin} value={locationField.display} rightElement={<EditActions onEdit={locationField.openEditor} />} />
                {locationField.modal}
            </div>
            
            {/* Ciudad de origen */}
            <div>
                <h3 className="font-semibold mb-3">Ciudad de origen</h3>
                <InfoRow icon={Home} value={originField.display} rightElement={<EditActions onEdit={originField.openEditor} />} />
                {originField.modal}
            </div>

            {/* Fecha de nacimiento */}
            <div>
                <h3 className="font-semibold mb-3">Fecha de nacimiento</h3>
                {(() => {
                    const display = birthday.value
                        ? birthday.value.split('-').reverse().join('/')
                        : <span className="text-muted-foreground">DD/MM/AAAA</span>;
                    return (
                        <InfoRow
                            icon={Cake}
                            value={display}
                            rightElement={<EditActions onEdit={birthday.openEditor} />}
                        />
                    );
                })()}
                {birthday.modal}
            </div>

            {/* Situación sentimental */}
            <div>
                <h3 className="font-semibold mb-3">Situación sentimental</h3>
                <InfoRow icon={Heart} value={relationship.value} rightElement={<EditActions onEdit={relationship.openEditor} />} />
                {relationship.modal}
            </div>

            {/* Familiares */}
            <div>
                <h3 className="font-semibold mb-3">Familiares</h3>
                <InfoRow icon={Users} value={family.value} rightElement={<EditActions onEdit={family.openEditor} />} />
                {family.modal}
            </div>

            {/* Género */}
            <div>
                <h3 className="font-semibold mb-3">Género</h3>
                <InfoRow 
                    icon={UserIcon} 
                    value={gender.value} 
                    subValue="Género"
                    rightElement={<EditActions onEdit={gender.openEditor} />}
                />
                {gender.modal}
            </div>

            {/* Pronombres */}
            <div>
                <h3 className="font-semibold mb-3">Pronombres</h3>
                <InfoRow 
                    icon={MessageSquare} 
                    value={pronouns.value} 
                    subValue="Pronombres del sistema"
                    rightElement={<EditActions onEdit={pronouns.openEditor} />} 
                />
                {pronouns.modal}
            </div>
             {/* Idiomas */}
             <div>
                <h3 className="font-semibold mb-3">Idiomas</h3>
                <InfoRow icon={Globe} value={languages.value} rightElement={<EditActions onEdit={languages.openEditor} />} />
                {languages.modal}
            </div>
        </div>
    );
}

export function EmploymentSection() {
    const field = useField("Experiencia laboral", "Empleo");
    return (
        <div className="animate-in fade-in">
            <h3 className="text-lg font-bold mb-4">Empleo</h3>
            <InfoRow icon={Briefcase} value={field.display} rightElement={<EditActions onEdit={field.openEditor} />} />
            {field.modal}
        </div>
    );
}

export function EducationSection() {
    const uni = useField("Universidad", "Universidad");
    const school = useField("Escuela secundaria", "Escuela secundaria");
    return (
        <div className="space-y-6 animate-in fade-in">
            <div>
                <h3 className="text-lg font-bold mb-4">Universidad</h3>
                <InfoRow icon={GraduationCap} value={uni.value} rightElement={<EditActions onEdit={uni.openEditor} />} />
                {uni.modal}
            </div>
            <div>
                <h3 className="text-lg font-bold mb-4">Escuela secundaria</h3>
                <InfoRow icon={Home} value={school.value} rightElement={<EditActions onEdit={school.openEditor} />} />
                {school.modal}
            </div>
        </div>
    );
}

export function HobbiesSection() {
    const field = useField("Pasatiempos", "Pasatiempos");
    return (
        <div className="animate-in fade-in">
            <h3 className="text-lg font-bold mb-4">Pasatiempos</h3>
            <InfoRow icon={Heart} value={field.display} rightElement={<EditActions onEdit={field.openEditor} />} />
            {field.modal}
            {/* Nota: puedes reemplazar el icono si prefieres otro (por ejemplo Gamepad2) */}
        </div>
    );
}

export function InterestsSection() {
    // for simplicity, allow editing the overall list description
    const field = useField("Música, TV, Películas, Juegos, Deportistas...", "Intereses");
    return (
        <div className="space-y-6 animate-in fade-in">
            <h3 className="font-semibold mb-3">Intereses</h3>
            <InfoRow icon={Music} value={field.display} rightElement={<EditActions onEdit={field.openEditor} />} />
            {field.modal}
        </div>
    );
}

export function TravelSection() {
    const field = useField("Lugares", "Viajes");
    return (
        <div className="animate-in fade-in">
            <h3 className="text-lg font-bold mb-1">Viajes</h3>
            <p className="text-sm text-muted-foreground mb-4">Agrega hasta 450 lugares que hayas visitado.</p>
            <InfoRow icon={Plane} value={field.display} rightElement={<EditActions onEdit={field.openEditor} />} />
            {field.modal}
        </div>
    );
}

export function LinksSection() {
    const field = useField("Sitios web, blogs, portfolios", "Enlaces");
    return (
        <div className="animate-in fade-in">
            <h3 className="text-lg font-bold mb-1">Enlaces</h3>
            <p className="text-sm text-muted-foreground mb-4">Agrega hasta 10 enlaces.</p>
            <InfoRow icon={LinkIcon} value={field.display} action rightElement={<EditActions onEdit={field.openEditor} />} />
            {field.modal}
        </div>
    );
}

export function ContactSection() {
    // reuse existing logic but convert to useField for phone/email too
    const phoneAField = useField("", "Teléfono A", "0414-1234567");
    const phoneBField = useField("", "Teléfono B", "0414-1234567");
    const emailField = useField("", "Correo electrónico", "usuario@ejemplo.com");

    return (
        <div className="space-y-6 animate-in fade-in">
            {/* Medios sociales */}
            <div>
                <h3 className="font-semibold mb-3">Medios sociales</h3>
                <InfoRow icon={LinkIcon} value="Medios sociales" action />
            </div>

            {/* Teléfono */}
            <div>
                <h3 className="font-semibold mb-3">Teléfono</h3>
                <div className="flex items-start gap-4 p-2 -ml-2">
                    <div className="w-9 h-9 mt-0.5 rounded-full bg-secondary/50 flex items-center justify-center shrink-0 text-muted-foreground">
                        <Phone className="w-5 h-5" />
                    </div>
                    <div className="flex-1">
                         <div className="flex justify-between items-start mb-2">
                            <div>
                                <p className="text-sm font-medium">{phoneAField.display}</p>
                                <p className="text-xs text-muted-foreground">Celular</p>
                            </div>
                            <div className="flex gap-2"><EditActions onEdit={phoneAField.openEditor} /></div>
                            {phoneAField.modal}
                         </div>
                         <div className="flex justify-between items-start">
                            <div>
                                <p className="text-sm font-medium">{phoneBField.display}</p>
                                <p className="text-xs text-muted-foreground">Celular</p>
                            </div>
                            <div className="flex gap-2"><EditActions onEdit={phoneBField.openEditor} /></div>
                            {phoneBField.modal}
                         </div>
                    </div>
                </div>
            </div>

            {/* Correo */}
            <div>
                <h3 className="font-semibold mb-3">Correo electrónico</h3>
                <InfoRow 
                    icon={Mail} 
                    value={emailField.display} 
                    subValue="Correo electrónico"
                    rightElement={<EditActions onEdit={emailField.openEditor} />}
                />
                {emailField.modal}
            </div>
        </div>
    );
}
export function NamesSection() {
    return (
        <div className="space-y-8 animate-in fade-in duration-300">
            {/* Bloque 1: Pronunciación */}
            <div>
                <h3 className="text-lg font-semibold mb-4 text-foreground">Pronunciación del nombre</h3>
                <div className="flex items-center gap-4 group cursor-pointer p-2 -ml-2 rounded-lg hover:bg-muted/50 transition-colors">
                    <div className="w-10 h-10 rounded-full bg-secondary flex items-center justify-center shrink-0">
                        <Mic className="w-5 h-5 text-foreground" />
                    </div>
                    <div>
                        <p className="text-sm font-medium text-blue-600 group-hover:underline">
                            Pronunciación del nombre
                        </p>
                        <p className="text-xs text-muted-foreground">
                            Añade una guía de audio para que sepan cómo decir tu nombre.
                        </p>
                    </div>
                </div>
            </div>

            {/* Separador visual si lo deseas */}
            {/* <hr className="border-border" /> */}

            {/* Bloque 2: Otros nombres */}
            <div>
                <h3 className="text-lg font-semibold mb-4 text-foreground">Otros nombres</h3>
                <div className="flex items-center gap-4 group cursor-pointer p-2 -ml-2 rounded-lg hover:bg-muted/50 transition-colors">
                    <div className="w-10 h-10 rounded-full bg-secondary flex items-center justify-center shrink-0">
                        <span className="text-lg font-serif font-bold text-foreground">Aa</span>
                    </div>
                    <div className="flex-1">
                        <p className="text-sm font-medium text-blue-600 group-hover:underline flex items-center gap-2">
                            <Plus className="w-3 h-3" /> Agregar otros nombres
                        </p>
                        <p className="text-xs text-muted-foreground">
                            Apodos, nombre de soltera, nombre alternativo...
                        </p>
                    </div>
                </div>
            </div>
        </div>
    );
}