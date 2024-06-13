import { Pencil1Icon, TrashIcon } from "@radix-ui/react-icons";
import { useDeleteTeacherMutation } from "../api/useDeleteTeacherMutation";
import { useState } from "react";
import EditModal from "../modals/EditModal.js";
import styles from "../../shared/style/table.module.css";
import { DeleteModal } from "../modals/DeleteModal.js";

export function EditCell(props) {
  const deleteTeacherMutation = useDeleteTeacherMutation();
  const id = props.props.original.id;
  const original = props.props.original;
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isDeleteModalOpen, setisDeleteModalOpen] = useState(false);

  const toggleEditModal = () => {
    setIsEditModalOpen(!isEditModalOpen);
  };
  function deleteTeacher() {
    setisDeleteModalOpen(!isDeleteModalOpen);
  }

  return (
    <div style={{ padding: "5px" }}>
      <div
        className={styles["btns-wrapper"]}
        onClick={() => toggleEditModal(id)}
      >
        <p>Изменить</p>
        <Pencil1Icon />
      </div>
      <div className={styles["btns-wrapper"]} onClick={deleteTeacher}>
        <p style={{ color: "var(--warning-color)" }}>Удалить</p>
        <TrashIcon />
      </div>
      {isEditModalOpen && (
        <EditModal
          original={original}
          toggleModal={toggleEditModal}
          isOpen={isEditModalOpen}
          id={id}
        />
      )}
      {isDeleteModalOpen && (
        <DeleteModal
          deleteTeacher={deleteTeacher}
          isOpen={isDeleteModalOpen}
          original={original}
        />
      )}
    </div>
  );
}
