import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useDeleteClassroomMutation } from "../api/useDeleteClassroomMutation";
import { useState } from "react";
import EditModal from "../modals/editModal.js";
import styles from "../../groups/table/table.module.css";

export function EditCell(props) {
  const deleteClassroomMutation = useDeleteClassroomMutation();
  const id = props.props.original.id;

  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteClassroom() {
    if (window.confirm("Вы уверены, что хотите удалить аудиторию?"))
    deleteClassroomMutation.mutateAsync(id);
  }

  return (
    <div style={{padding: '5px'}}>
      <div
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </div>
      <div className={styles["btns-wrapper"]} onClick={deleteClassroom}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </div>
      {isEditModalOpen && (
        <EditModal
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          id={id}
        />
      )}
    </div>
  );
}
